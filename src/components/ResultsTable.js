import React, { useEffect, useState } from 'react'
import ReactHtmlParser from 'react-html-parser'; 
import { useNavigate } from "react-router-dom";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableFooter from '@mui/material/TableFooter';
import TablePagination from '@mui/material/TablePagination';
import Paper from '@mui/material/Paper';
import IconButton from '@mui/material/IconButton';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';
import { useTheme } from '@mui/material/styles';
import Box from '@mui/material/Box';
import PropTypes from 'prop-types';
import axios from '../axios';
import { useParams } from "react-router-dom";
import ResultsQuery from './ResultsQuery';


function TablePaginationActions(props) {
    const theme = useTheme();
    const { count, page, rowsPerPage, onPageChange } = props;
  
    const handleFirstPageButtonClick = (event) => {
      onPageChange(event, 0);
    };
  
    const handleBackButtonClick = (event) => {
      onPageChange(event, page - 1);
    };
  
    const handleNextButtonClick = (event) => {
      onPageChange(event, page + 1);
    };
  
    const handleLastPageButtonClick = (event) => {
      onPageChange(event, Math.max(0, Math.ceil(count / rowsPerPage) - 1));
    };
  
    return (
      <Box sx={{ flexShrink: 0, ml: 2.5 }}>
        <IconButton
          onClick={handleFirstPageButtonClick}
          disabled={page === 0}
          aria-label="first page"
        >
          {theme.direction === 'rtl' ? <LastPageIcon /> : <FirstPageIcon />}
        </IconButton>
        <IconButton
          onClick={handleBackButtonClick}
          disabled={page === 0}
          aria-label="previous page"
        >
          {theme.direction === 'rtl' ? <KeyboardArrowRight /> : <KeyboardArrowLeft />}
        </IconButton>
        <IconButton
          onClick={handleNextButtonClick}
          disabled={page >= Math.ceil(count / rowsPerPage) - 1}
          aria-label="next page"
        >
          {theme.direction === 'rtl' ? <KeyboardArrowLeft /> : <KeyboardArrowRight />}
        </IconButton>
        <IconButton
          onClick={handleLastPageButtonClick}
          disabled={page >= Math.ceil(count / rowsPerPage) - 1}
          aria-label="last page"
        >
          {theme.direction === 'rtl' ? <FirstPageIcon /> : <LastPageIcon />}
        </IconButton>
      </Box>
    );
  }
  
  TablePaginationActions.propTypes = {
    count: PropTypes.number.isRequired,
    onPageChange: PropTypes.func.isRequired,
    page: PropTypes.number.isRequired,
    rowsPerPage: PropTypes.number.isRequired,
  };

function ResultsTable(props) {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);
    const [results, setResults] = useState([])
    const { query } = useParams();
    const [hasSearched, setHasSearched] = useState(0);
    
    useEffect(() => {
        async function search() {
            setHasSearched(1)
            console.log("NEW SEARCH!!")
            await axios.get(`/search?q=${query}`).then((response) => {
                setResults(response.data)
            });
        }
        search()
    }, [query])

    

    let navigate = useNavigate();
    async function readMore(query, index) {
        const readUrl = '/read/' + query + '/' + index;
        navigate(readUrl);
    }
    const handleChangePage = (event, newPage) => {
        setPage(newPage);
      };
    
      const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
      };

    return (
      <div>
        <h1 className="title">Shakesearch</h1>
        <div className="resultsBody">
          <ResultsQuery/>
          <div className="tableDiv">
              {hasSearched && <p> Displaying {results.length} Results for "{query}"</p>} 
              {hasSearched && results.length === 0 && <p> No searches found. Try again</p>} 
              {hasSearched && results.length > 0 && <TableContainer component={Paper}>
              <Table className="resultsTable" aria-label="simple table">
                  <TableHead>
                  <TableRow>
                      <TableCell sx={{border: "1px", fontWeight: "550" }}>Text Preview</TableCell>
                      <TableCell sx={{borderBottom: "1px", fontWeight: "550"}}>Character Position</TableCell>
                  </TableRow>
                  </TableHead>
                  <TableBody>
                      {(rowsPerPage > 0
                              ? results.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                              : results
                          ).map((result) => (
                          <TableRow
                          key={result.Index}
                          >
                          <TableCell sx={{ fontFamily: "Courier New", fontWeight: "550", border: "1px" }} align="left">{ReactHtmlParser(result.Result)}</TableCell>
                          <TableCell  sx={{ fontFamily: "Courier New", fontWeight: "500", borderBottom: "1px" }} component="th" scope="row">
                              {<span className="readMore" onClick={() => readMore(result.Query, result.Index)}>{result.Index}</span>}
                          </TableCell>
                          </TableRow>
                      ))}
                  
                  </TableBody>
                  <TableFooter>
                      <TableRow>
                          <TablePagination
                          sx={{ borderBottom: "1px" }}
                          rowsPerPageOptions={[5, 10, 25, { label: 'All', value: -1 }]}
                          colSpan={2}
                          count={results.length}
                          rowsPerPage={rowsPerPage}
                          page={page}
                          SelectProps={{
                              inputProps: {
                              'aria-label': 'rows per page',
                              },
                              native: true,
                          }}
                          onPageChange={handleChangePage}
                          onRowsPerPageChange={handleChangeRowsPerPage}
                          ActionsComponent={TablePaginationActions}
                          />
                      </TableRow>
                  </TableFooter>
              </Table>
              </TableContainer>}
          </div>
        </div>
      </div>
    );

}

export default ResultsTable
