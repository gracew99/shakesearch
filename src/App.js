import './App.css';
import Query from './components/Query';
import Table from './components/Table';
import React, { useState } from 'react';

function App() {
  const [results, setResults] = useState([]);

  return (
    <div className="App">
      <header className="App-header">
      </header>
      <Query setResults={setResults}></Query>
      <Table results={results}></Table>
    </div>
  );
}

export default App;
