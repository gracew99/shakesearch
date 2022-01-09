import React, { useEffect, useState } from 'react'
import axios from '../axios';
import { useParams } from "react-router-dom";
import ReactHtmlParser from 'react-html-parser'; 
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';
import ArrowForwardIosIcon from '@mui/icons-material/ArrowForwardIos';

function ReadPage() {
    const { query, index } = useParams();
    const [text, setText] = useState();
    let [anchorIndex, setAnchorIndex] = useState(parseInt(index));
    
    async function read(anchorIndex) {
        const highlight = parseInt(anchorIndex) === parseInt(index);
        console.log("requesting", anchorIndex, highlight); 
        const axiosUrl = "/read?query=" + query + "&index=" + anchorIndex + "&highlight=" + highlight;
        await axios.get(axiosUrl).then((response) => {
            setText(ReactHtmlParser(response.data))
        });
    }
    useEffect(() => {
        read(index)
    }, [])
    return (
        <div>
            <h1 className="title">Shakesearch</h1>
            <div className="readPage">
                <div onClick={() => {read(anchorIndex-1000); setAnchorIndex(currIndex => currIndex-1000)}}>
                    <ArrowBackIosIcon></ArrowBackIosIcon>
                </div>
                <div style={{marginBottom: "10%"}}>
                    {text}
                </div>                
                <div onClick={() => {read(anchorIndex+1000); setAnchorIndex(currIndex => currIndex+1000)}}>
                    <ArrowForwardIosIcon></ArrowForwardIosIcon>
                </div>
            </div>
        </div>
    )
}

export default ReadPage
