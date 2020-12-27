import React, {Component} from "react"
import "./ChatHistory.scss"

class ChatHistory extends Component {
    render(){
        const messages = this.props.chatHistory.map((msg, idx) => (
            <p key={idx}>{msg.data}</p>
        ));

        return (
            <div className="ChatHistory">
                <h2>Chat history</h2>
                {messages}
            </div>
        );
    }
}

export default ChatHistory