import React, { Component } from 'react';
import Header from './components/Header/Header';
import ChatContainer from './components/ChatContainer';
import './App.css';

class App extends Component {
 

  render() {
    return (
      <div className="App">
        <Header />
        <ChatContainer/>
      </div>
    );
  }
}

export default App;