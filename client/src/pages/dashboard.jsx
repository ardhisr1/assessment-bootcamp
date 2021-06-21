import React from 'react'
import { useSelector } from 'react-redux'


export default function Dashboard() {
    const token = useSelector(state => state.userLogin)

    return (
        <div>
            <h1>Dashboard</h1>
            <h2>token :{token.token}</h2>
        </div>
    )
}
