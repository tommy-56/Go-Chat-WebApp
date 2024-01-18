import React, { Component } from "react";
import "./Message.css";

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    };
  }

  render() {
    return(
       <div className="Message">
          {this.state.message.readableName ? (
            <div>
              {/* Render content when readableName is not empty */}
              {this.state.message.readableName +" says: "}{this.state.message.body}
            </div>
          ) : (
            <div>
              {/* Render content when readableName is empty */}
              {this.state.message.body}
            </div>
        )}
      </div>
    );
  }
}

export default Message;