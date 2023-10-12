import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { Box, Button, Grid, IconButton, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material"
import NodeIcon from "../NodeIcon"
import { faTimes } from "@fortawesome/free-solid-svg-icons"
import { AssetGroupChangelogEntry } from "../AssetGroupAutocomplete"
import { FC } from "react"

const AssetGroupChangelogTable: FC<{
    addRows: AssetGroupChangelogEntry[],
    removeRows: AssetGroupChangelogEntry[],
    onRemove: (entry: AssetGroupChangelogEntry) => void,
    onCancel: () => void,
    onSubmit: () => void,
}> = ({
    addRows,
    removeRows,
    onRemove,
    onCancel,
    onSubmit,
}) => {
    return (
        <>
            <TableContainer>
                <Table size='small'>
                    {addRows.length > 0 && (
                        <AssetGroupChangelogRows
                            title="Add to Group"
                            rows={addRows}
                            onRemove={onRemove}
                        />
                    )}
                    {removeRows.length > 0 && (
                        <AssetGroupChangelogRows
                            title="Remove from Group"
                            rows={removeRows}
                            onRemove={onRemove}
                        />
                    )}
                </Table>
            </TableContainer>
            <Box mt={1}>
                <Grid container direction='row' justifyContent='flex-end' spacing={1}>
                    <Grid item>
                        <Button color='inherit' size='small' onClick={onCancel}>
                            Cancel
                        </Button>
                    </Grid>
                    <Grid item>
                        <Button
                            size='small'
                            color='primary'
                            variant='contained'
                            disableElevation
                            onClick={onSubmit}>
                            Confirm Changes
                        </Button>
                    </Grid>
                </Grid>
            </Box>
        </>
    )
}

const AssetGroupChangelogRows: FC<{
    title: string,
    rows: AssetGroupChangelogEntry[],
    onRemove: (entry: AssetGroupChangelogEntry) => void,
}> = ({ title, rows, onRemove }) => {
    return (
        <>
            <TableHead>
                <TableRow>
                    <TableCell colSpan={2}>{title}</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {rows.map((row) => (
                    <TableRow key={row.objectid}>
                        <TableCell padding='none'>
                            <IconButton
                                size='small'
                                onClick={() => onRemove(row)}
                            >
                                <FontAwesomeIcon icon={faTimes} />
                            </IconButton>
                        </TableCell>
                        <TableCell
                            style={{
                                whiteSpace: 'nowrap',
                            }}>
                            <NodeIcon nodeType={row.type} />
                            {row.name}
                            <br />
                            <small>{row.objectid}</small>
                        </TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </>
    )
}

export default AssetGroupChangelogTable;