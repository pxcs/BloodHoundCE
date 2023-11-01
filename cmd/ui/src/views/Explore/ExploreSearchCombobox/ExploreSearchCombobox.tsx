// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { List, ListItem, ListItemText, Paper, TextField, useTheme } from '@mui/material';
import { useCombobox } from 'downshift';
import { useEffect, useState } from 'react';
import { NodeIcon, SearchResultItem } from 'bh-shared-ui';
import { getEmptyResultsText, getKeywordAndTypeValues, SearchResult, useSearch } from 'src/hooks/useSearch';
import { AppState, useAppDispatch } from 'src/store';
import {
    destinationNodeSelected,
    setSearchValue,
    sourceNodeSelected,
    startSearchAction,
} from 'src/ducks/searchbar/actions';
import { PRIMARY_SEARCH, SEARCH_TYPE_EXACT, PATHFINDING_SEARCH, SearchNodeType } from 'src/ducks/searchbar/types';
import { useSelector } from 'react-redux';

const ExploreSearchCombobox: React.FC<{
    labelText: string;
    disabled?: boolean;
    searchType: typeof PRIMARY_SEARCH | typeof PATHFINDING_SEARCH;
}> = ({ labelText, disabled = false, searchType }) => {
    const theme = useTheme();
    const dispatch = useAppDispatch();

    const [showInputIcon, setShowInputIcon] = useState(true);

    const { primary, secondary } = useSelector((state: AppState) => state.search);
    const inputValue = searchType === PRIMARY_SEARCH ? primary.searchTerm : secondary.searchTerm;
    const selectedItem = searchType === PRIMARY_SEARCH ? primary.value : secondary.value;
    const controlledOpenMenu = searchType === PRIMARY_SEARCH ? primary.openMenu : secondary.openMenu;

    const { keyword, type } = getKeywordAndTypeValues(inputValue);
    const { data, error, isError, isLoading, isFetching } = useSearch(keyword, type);

    const { isOpen, getMenuProps, getInputProps, getComboboxProps, highlightedIndex, getItemProps, openMenu } =
        useCombobox({
            items: data || [],
            onInputValueChange: ({ type, inputValue }) => {
                if (
                    inputValue !== undefined &&
                    type !== useCombobox.stateChangeTypes.ControlledPropUpdatedSelectedItem
                ) {
                    dispatch(startSearchAction(inputValue, searchType));
                }
            },
            inputValue,
            selectedItem,
            onSelectedItemChange: ({ type, selectedItem }) => {
                if (selectedItem) {
                    dispatch(setSearchValue(selectedItem as SearchNodeType, searchType, SEARCH_TYPE_EXACT));
                    // todo: need to dispatch action that the source/end node was selected
                    if (searchType === PRIMARY_SEARCH) {
                        dispatch(sourceNodeSelected());
                    } else if (searchType === PATHFINDING_SEARCH) {
                        dispatch(destinationNodeSelected());
                    }
                }
            },
            itemToString: (item) => (item ? item.name || item.objectid : ''),
            onStateChange: ({ type, inputValue, selectedItem }) => {
                // remove icon when input is empty
                if (type === useCombobox.stateChangeTypes.InputChange) {
                    if (!inputValue) {
                        setShowInputIcon(false);
                    }
                }

                // show icon when item is selected via click or enter key
                if (type === useCombobox.stateChangeTypes.ItemClick) {
                    setShowInputIcon(true);
                }
                if (type === useCombobox.stateChangeTypes.InputKeyDownEnter) {
                    setShowInputIcon(true);
                }

                // show icon when input is controlled via props
                if (type === useCombobox.stateChangeTypes.ControlledPropUpdatedSelectedItem) {
                    setShowInputIcon(true);
                }
            },
        });

    // handy when combobox need's to be opened from outside the component, e.g. the <ContextMenu /> can set the input value
    useEffect(() => {
        if (controlledOpenMenu) {
            openMenu();
        }
    }, [controlledOpenMenu, openMenu]);

    const disabledText: string = getEmptyResultsText(
        isLoading,
        isFetching,
        isError,
        error,
        inputValue,
        type,
        keyword,
        data
    );

    return (
        <div {...getComboboxProps()} style={{ position: 'relative' }}>
            <TextField
                placeholder={labelText}
                variant='outlined'
                size='small'
                fullWidth
                disabled={disabled}
                inputProps={{
                    'aria-label': labelText,
                }}
                InputProps={{
                    style: {
                        backgroundColor: disabled ? theme.palette.action.disabled : 'inherit',
                        fontSize: theme.typography.pxToRem(14),
                    },
                    startAdornment: showInputIcon && selectedItem?.type && <NodeIcon nodeType={selectedItem?.type} />,
                }}
                {...getInputProps({ onFocus: openMenu, refKey: 'inputRef' })}
                data-testid='explore_search_input-search'
            />
            <div
                style={{
                    position: 'absolute',
                    marginTop: '1rem',
                    zIndex: 1300,
                }}>
                <Paper style={{ display: isOpen ? 'inherit' : 'none' }}>
                    <List
                        {...getMenuProps()}
                        dense
                        style={{
                            width: '100%',
                        }}
                        data-testid='explore_search_result-list'>
                        {disabledText ? (
                            <ListItem disabled dense>
                                <ListItemText primary={disabledText} />
                            </ListItem>
                        ) : (
                            data!.map((item: SearchResult, index: any) => {
                                return (
                                    <SearchResultItem
                                        item={{
                                            label: item.name,
                                            objectId: item.objectid,
                                            kind: item.type,
                                        }}
                                        index={index}
                                        key={index}
                                        highlightedIndex={highlightedIndex}
                                        keyword={keyword}
                                        getItemProps={getItemProps}
                                    />
                                );
                            })
                        )}
                    </List>
                </Paper>
            </div>
        </div>
    );
};

export default ExploreSearchCombobox;
