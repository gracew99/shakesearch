import React, { useState } from 'react';
import axios from '../axios';
function Query(props) {
    const [query, setQuery] = useState("");
    function handleChange(event) {
      setQuery(event.target.value);
    }
     async function search(event) {
       event.preventDefault();
        console.log(query)
        await axios.get(`/search?q=${query}`).then((response) => {
          props.setResults(response.data.Results)
        });
        setQuery("")
      }
    return (
        <div style={{display: "flex", justifyContent: "center", alignItems: "center"}}>
              <form id="form" onSubmit={search}>
                  <input className="input" type="text" id="query" name="query" placeholder="Search..."onChange={handleChange}/>
                  <button type="submit">Search</button>
              </form>
        </div>
    )
}

export default Query
