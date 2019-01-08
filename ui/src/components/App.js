import React, { Component } from "react";
import TodoList from "./TodoList";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      articles: []
    };
  }
  componentDidMount() {
    let myHeaders = new Headers();

    let myInit = {
      method: "GET",
      headers: myHeaders,
      mode: "cors",
      cache: "default"
    };

    fetch("http://localhost:10000/sel", myInit)
      .then(res => res.json()) //response type
      .then(data => {
        console.log(data); //log the data
        this.setState({ articles: data });
      });
  }

  render() {
    return (
      <div className="eleven wide column">
        <TodoList todos={this.state.articles} />
      </div>
    );
  }
}

export default App;
