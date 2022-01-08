import React, { useState } from 'react';
import axios from '../axios';
import TextField from '@mui/material/TextField';
import SearchIcon from '@mui/icons-material/Search';
import InputAdornment from '@mui/material/InputAdornment';
import IconButton from '@mui/material/IconButton';

function Query(props) {
    const [query, setQuery] = useState(""); // local query
    function handleChange(event) {
      setQuery(event.target.value);
    }
     async function search(event) {
       event.preventDefault();
        console.log(query)
        await axios.get(`/search?q=${query}`).then((response) => {
          props.setResults(response.data)
          props.setQuery(query) // announce to higher
        });
        setQuery("")
      }
    return (
        <div style={{marginTop: "10%", display: "flex", justifyContent: "center", alignItems: "center"}}>
              <form id="form" style={{flex: 0.7}} onSubmit={search}>
                  <TextField className="input" id="standard-basic" variant="standard" name="query" placeholder="Search..." onChange={handleChange}
                    InputProps={{
                      endAdornment: (
                        <InputAdornment>
                          <IconButton>
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

export default Query
