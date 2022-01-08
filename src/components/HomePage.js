import Query from './Query';
import ResultsTable from './ResultsTable';
import React, { useState } from 'react';

function HomePage() {
    const [results, setResults] = useState([]);
    const [queryLength, setQueryLength] = useState([]);
    return (
        <div>
            <Query setResults={setResults} setQueryLength={setQueryLength}></Query>
            <ResultsTable results={results} queryLength={queryLength}></ResultsTable>
        </div>
    )
}

export default HomePage
