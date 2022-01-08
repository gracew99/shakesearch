import Query from './Query';
import ResultsTable from './ResultsTable';
import React, { useState } from 'react';

function HomePage() {
    const [results, setResults] = useState([]);
    const [query, setQuery] = useState(""); // higher query
    return (
        <div>
            <Query setResults={setResults} setQuery={setQuery}></Query>
            <ResultsTable results={results} query={query}></ResultsTable>
        </div>
    )
}

export default HomePage
