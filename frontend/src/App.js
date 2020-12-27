import React, {Component} from "react";
// import logo from './logo.svg';
import './App.css';
import {connect, sendMsg} from "./api";

import Header from "./components/Header";
import ChatHistory from "./components/ChatHistory";

class App extends Component{
  constructor(props){
    super(props);

    //save status
    this.state = {
        //chat history in status
        chatHistory: []
    }

    // connect(); //before relocated
  }

  componentDidMount() {
      connect((msg) => {
          console.log("New message arrived")
          this.setState(prevState => ({
              chatHistory: [...this.state.chatHistory, msg]
          }))
          console.log(this.state)
      });
  }

    send(){
    console.log("hello")
    sendMsg("hello")
  }

  render(){
    return (
        <div className="App">
            <Header></Header>
            <ChatHistory chatHistory={this.state.chatHistory}></ChatHistory>
          <button onClick={this.send}>HIT HELLO</button>
        </div>
    )
  }
}

export default App;
