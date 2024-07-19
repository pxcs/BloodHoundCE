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

import { Theme } from '@mui/material';

export const NODE_GRAPH_RENDER_LIMIT = 1000;

export const ZERO_VALUE_API_DATE = '0001-01-01T00:00:00Z';

export const TIER_ZERO_TAG = 'admin_tier_0';
export const TIER_ZERO_LABEL = 'High Value';

export const lightPalette = {
    primary: {
        main: '#33318F',
        dark: '#261F7A',
    },
    secondary: {
        main: '#1A30FF',
        dark: '#0524F0',
    },
    tertiary: {
        main: '#5CC791',
        dark: '#02C577',
    },
    color: {
        primary: '#1D1B20',
        links: '#1A30FF',
        error: '#B44641',
    },
    neutral: {
        primary: '#FFFFFF',
        secondary: '#F4F4F4',
        tertiary: '#E3E7EA',
        quaternary: '#DADEE1',
        quinary: '#CACFD3',
    },
    background: {
        paper: '#fafafa',
        default: '#e4e9eb',
    },
    low: 'rgb(255, 195, 15)',
    moderate: 'rgb(255, 97, 66)',
    high: 'rgb(205, 0, 117)',
    critical: 'rgb(76, 29, 143)',
};

export const darkPalette = {
    primary: {
        main: '#33318F',
        dark: '#261F7A',
    },
    secondary: {
        main: '#1A30FF',
        dark: '#0524F0',
    },
    tertiary: {
        main: '#5CC791',
        dark: '#02C577',
    },
    color: {
        primary: '#FFFFFF',
        links: '#99A3FF',
        error: '#E9827C',
    },
    neutral: {
        primary: '#121212',
        secondary: '#222222',
        tertiary: '#272727',
        quaternary: '#2C2C2C',
        quinary: '#2E2E2E',
    },
    background: {
        paper: '#211F26',
        default: '#121212',
    },
    low: 'rgb(255, 195, 15)',
    moderate: 'rgb(255, 97, 66)',
    high: 'rgb(205, 0, 117)',
    critical: 'rgb(76, 29, 143)',
};

export const typography: Partial<Theme['typography']> = {
    h1: {
        fontWeight: 400,
        fontSize: '1.8rem',
        lineHeight: 2,
        letterSpacing: 0,
    },
    h2: {
        fontWeight: 500,
        fontSize: '1.5rem',
        lineHeight: 1.5,
        letterSpacing: 0,
    },
    h3: {
        fontWeight: 500,
        fontSize: '1.2rem',
        lineHeight: 1.25,
        letterSpacing: 0,
    },
    h4: {
        fontWeight: 500,
        fontSize: '1.25rem',
        lineHeight: 1.5,
        letterSpacing: 0,
    },
    h5: {
        fontWeight: 700,
        fontSize: '1.125rem',
        lineHeight: 1.5,
        letterSpacing: 0.25,
    },
    h6: {
        fontWeight: 700,
        fontSize: '1.0rem',
        lineHeight: 1.5,
        letterSpacing: 0.25,
    },
};

export const components = (palette: typeof darkPalette): Partial<Theme['components']> => ({
    MuiButton: {
        styleOverrides: {
            root: {
                borderRadius: 999, // capsule-shaped buttons
            },
        },
    },
    MuiAccordionSummary: {
        styleOverrides: {
            root: {
                flexDirection: 'row-reverse',
            },
            content: {
                marginRight: '4px',
            },
        },
    },
    MuiLink: {
        styleOverrides: {
            root: {
                color: palette.color.links,
            },
        },
    },
    MuiInputLabel: {
        styleOverrides: {
            root: {
                '&.Mui-focused': {
                    color: palette.color.links,
                },
            },
        },
    },
    MuiTextField: {
        styleOverrides: {
            root: {
                '&:hover .MuiInputBase-root .MuiOutlinedInput-notchedOutline': {
                    borderColor: palette.color.links,
                },
                '& .MuiInputBase-root.Mui-focused .MuiOutlinedInput-notchedOutline': {
                    borderColor: palette.color.links,
                }
            },
        }
    },
    MuiInput: {
        styleOverrides: {
            underline: {
                '&:after': {
                    borderBottom: `2px solid ${palette.color.links}`,
                },
                '&:hover:not($disabled):not($focused):not($error):before': {
                    borderBottom: `2px solid ${palette.color.links}`,
                },
            },
        },
    },
    MuiDialog: {
        styleOverrides: {
            root: {
                '& .MuiPaper-root': {
                    backgroundImage: 'unset',
                },
            },
        },
    },
    MuiPopover: {
        styleOverrides: {
            root: {
                '& .MuiPaper-root': {
                    backgroundImage: 'unset',
                },
            },
        },
    },
    MuiCheckbox: {
        styleOverrides: {
            root: {
                '& svg': {
                    color: palette.color.primary,
                },
            },
        },
    },
    MuiTabs: {
        styleOverrides: {
            root: {
                '& .MuiTab-labelIcon': {
                    color: palette.color.links,
                },
                '& .MuiTab-labelIcon.Mui-selected': {
                    color: palette.color.links,
                },
                '& .MuiTab-labelIcon:not(.Mui-selected)': {
                    color: palette.color.primary,
                },
                '& .MuiTabs-indicator': {
                    backgroundColor: palette.color.links,
                },
            },
        },
    },
});
