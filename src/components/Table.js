import React from 'react'

function Table(props) {

    return (
        <div className="tableDiv">
            <table id="table">
            <tbody id="table-body">
                {
                    props.results && props.results.map(result => {
                        return (
                            <tr><p>{result}</p></tr> 
                        )
                })}
            </tbody>
            </table>
        </div>
    )
}

export default Table
