import HomeQuery from './HomeQuery';
import React from 'react';

function HomePage() {
    return (
        <div className="home">
            <h1 className="title">Shakesearch</h1>
            <div className="homeBody">
                <HomeQuery />
            </div>
            <div className="homeImage">
                <img className="shakespeare" src='/shakespeare.png' alt="shakespeare"></img>
            </div>
        </div>
    )
}

export default HomePage
