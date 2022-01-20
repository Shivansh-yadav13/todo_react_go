import React, { useState, useEffect } from "react";
import axios from "axios";
import "./App.css";

function App() {
  const [todoList, setTodoList] = useState(null);

  const getAllTodos = async () => {
    await axios.get("http://localhost:3001/api/get_todos").then((res) => {
      setTodoList(res.data);
    });
  };
  useEffect(() => {
    getAllTodos();
  }, []);
  return (
    <div className="App">
      <h1>Todo - List</h1>
      <div className="list-complete">
        {todoList ? (
          todoList.map((todo, key) => (
            <div className="todo-item" key={key}>
              <h2>{todo.title}</h2>
              <p>{todo.description}</p>
              <li>{todo.status}</li>
            </div>
          ))
        ) : (
          <h1>Loading...</h1>
        )}
      </div>
      <div className="list-incomplete">
          <li>NOOB</li>
      </div>
    </div>
  );
}

export default App;
