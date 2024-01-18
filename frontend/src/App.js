import React, { Component } from 'react';
import { Routes, Route } from "react-router-dom";
import Login from './components/Login';
import ChatContainer from './components/ChatContainer';
import './App.css';
import { connect, sendMsg } from './api';

class App extends Component {
  

  render() {
    return (
      <div className="App">
        <Routes>
        <Route path="/" element={ <ChatContainer/> } />
        <Route path="login" element={<Login/>} />
      </Routes>
      </div>
    );
  }
}

export default App;