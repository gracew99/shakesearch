import HomeQuery from './HomeQuery';
import React from 'react';

function HomePage() {
    return (
        <div className="home">
            <h1 className="title">Shakesearch</h1>
            <div className="homeBody">
                <HomeQuery />
            </div>
        </div>
    )
}

export default HomePage
