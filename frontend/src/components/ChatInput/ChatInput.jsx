import React, {Component} from "react";
import "./ChatInput.scss";

// this.props.send is set by outter App.js
// like this    <ChatInput send={this.send()}></ChatInput>

class ChatInput extends Component{
    render(){
        return (
            <div className="ChatInput">
                <input onKeyDown={this.props.send}/>
            </div>
        )
    }
}

export default ChatInput;