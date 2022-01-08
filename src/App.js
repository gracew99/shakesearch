import './App.css';
import React from 'react';
import { Routes, Route } from "react-router-dom";
import HomePage from './components/HomePage';
import ReadPage from './components/ReadPage';
import ResultsTable from './components/ResultsTable';

function App() {

  return (
    <div className="App">
      <header className="App-header">
      </header>
        <Routes>
          <Route exact path="/" element={<HomePage/>} />
          <Route path="/results/:query" element={<ResultsTable/>} />
          <Route path="/read/:query/:index" element={<ReadPage/>} />
        </Routes>

    </div>
  );
}

export default App;
