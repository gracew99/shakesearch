import React, { useEffect, useState } from 'react'
import axios from '../axios';

function ReadPage() {
    const [text, setText] = useState([]);
    async function read() {
        await axios.get(`/read`).then((response) => {
            setText(response.data.substring(0,5000))
        });
    }
    useEffect(() => {
        read()
    }, [])
    return (
        <div>
        
            {text}
        </div>
    )
}

export default ReadPage
