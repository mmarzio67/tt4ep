import "./TodoList.css";
import TodoItem from "./TodoItem";
import React from "react";

const TodoList = ({ todos, onTodoSelect }) => {
  const renderedList = todos.map(todo => {
    return (
      <TodoItem
        key={todo.descr}
        onTodoSelect={onTodoSelect}
        todo={todo}
      />
    );
  });
  return <div className="ui relaxed divided list">{renderedList}</div>;
};

export default TodoList;
