import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import SearchIcon from '@mui/icons-material/Search';
import InputAdornment from '@mui/material/InputAdornment';
import IconButton from '@mui/material/IconButton';
import { useNavigate } from "react-router-dom";

function ResultsQuery(props) {
    const [query, setQuery] = useState(""); 
    let navigate = useNavigate();

    function handleChange(event) {
      setQuery(event.target.value);
    }
     async function search(event) {
        event.preventDefault();

        const resultsUrl = '/results/' + query;
        navigate(resultsUrl);
        console.log(query)
        setQuery("")
      }
    return (
        <div style={{marginTop: "10%", display: "flex", justifyContent: "center", alignItems: "center"}}>
              <form id="form" style={{flex: 0.7}} onSubmit={search}>
                  <TextField className="input" id="standard-basic" variant="standard" name="query" placeholder="Search..." onChange={handleChange}
                    InputProps={{
                      endAdornment: (
                        <InputAdornment>
                          <IconButton onClick={search}>
                            <SearchIcon />
                          </IconButton>
                        </InputAdornment>
                      )
                    }}
                  />
              </form>
        </div>
    )
}

export default ResultsQuery
