import React from "react"
import "./App.css"
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Register from "./pages/register";
import Login from "./pages/login"
import Home from "./pages/home"
import Dashboard from "./pages/dashboard"

function App() {
  return (
    <div>
      <Router>
        <Switch>
          <Route path={"/register"}>
            <Register/>
          </Route>
          <Route path={"/login"}>
            <Login/>
          </Route>
          <Route path={"/dashboard"}>
            <Dashboard/>
          </Route>
          <Route path={"/"}>
            <Home/>
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
