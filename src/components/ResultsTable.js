import React from 'react'
import ReactHtmlParser from 'react-html-parser'; 
import { useNavigate } from "react-router-dom";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

function ResultsTable(props) {

    let navigate = useNavigate();
    async function readMore(query, index) {
        const readUrl = '/read/' + query + '/' + index;
        navigate(readUrl);
    }

    return (
        <div className="tableDiv">
            {props.results && props.results.length > 0 && <p> Displaying {props.results.length} Results for "{props.query}"</p>} 
            {props.results && props.results.length > 0 && <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                <TableRow>
                    <TableCell>Text</TableCell>
                    <TableCell>Character Position</TableCell>
                </TableRow>
                </TableHead>
                {props.results.map((result) => (
                    <TableRow
                    key={result.Index}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    >
                    <TableCell align="left">{ReactHtmlParser(result.Result)}</TableCell>
                    <TableCell component="th" scope="row">
                        {<span className="readMore" onClick={() => readMore(props.query, result.Index)}>{result.Index}</span>}
                    </TableCell>
                    </TableRow>
                ))}
                <TableBody>
                
                </TableBody>
            </Table>
            </TableContainer>}
        </div>
    );

}

export default ResultsTable
