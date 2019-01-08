import "./TodoItem.css";
import React from "react";

const TodoItem = ({ todo, onTodoSelect }) => {
  // above: instead of props (witch contains the video object from the API)
  // I just destructure and take only the video portion ({video})
  return (
    <div onClick={() => onTodoSelect(todo)} className="todo-item item">
      <div className="content">
        <div className="header">Timestamp: {todo.timestamp}</div>
        <div className="header">Description: {todo.descr}</div>
        <div className="header">Action: {todo.action}</div>
        <div className="header">Project: {todo.project}</div>
      </div>
    </div>
  );
};

export default TodoItem;
