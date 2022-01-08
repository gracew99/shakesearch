import './App.css';
import React from 'react';
import { Routes, Route } from "react-router-dom";
import HomePage from './components/HomePage';
import ReadPage from './components/ReadPage';

function App() {

  return (
    <div className="App">
      <header className="App-header">
      </header>
        <Routes>
          <Route exact path="/" element={<HomePage/>} />
          <Route exact path="/read" element={<ReadPage/>} />
        </Routes>

    </div>
  );
}

export default App;
