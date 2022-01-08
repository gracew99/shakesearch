import React, { useEffect, useState } from 'react'
import axios from '../axios';
import { useParams } from "react-router-dom";
import ReactHtmlParser from 'react-html-parser'; 

function ReadPage() {
    const { query, index } = useParams();
    const [text, setText] = useState();
    
    useEffect(() => {
        async function read() {
            const axiosUrl = "/read?query=" + query + "&index=" + index;
            await axios.get(axiosUrl).then((response) => {
                setText(ReactHtmlParser(response.data))
            });
        }
        read()
    }, [])
    return (
        <div style={{display: "flex", marginTop: "10%", alignItems: "center", justifyContent: "center"}}>
            <div>
                {text}
            </div>
        </div>
    )
}

export default ReadPage
