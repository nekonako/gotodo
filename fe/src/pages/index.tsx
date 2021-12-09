import axios from 'axios';
import { useRef, useState } from 'react';
import { API_URL } from '../constants';
import { getTodos, Todo, useGetTodos } from './hooks/useGetTodos';

export default function Index() {
  let { todos, setTodos } = useGetTodos();
  const [todoUpdate, setTodoUpdate] = useState({
    id: null,
  });
  const [todo, setTodo] = useState('');
  const dataAddTodo = useRef(null);
  const [addTodo, setAddTodo] = useState(false);
  const searchTodo = useRef(null);

  if (todos === null) return <div>loading</div>;

  const updateToggle = (id: number) => {
    setTodoUpdate({ id: todoUpdate.id !== null ? null : id });
  };

  const edit = async (id: number) => {
    try {
      const res = await axios.put(`${API_URL}/update-todo`, {
        id: id,
        todo: todo,
      });
      if (res.data.code == 200) {
        const newTodos = await getTodos();
        setTodos(newTodos);
        updateToggle(id);
      }
    } catch (err) {
      console.log(err);
    }
  };

  const remove = async (id: number) => {
    try {
      const res = await axios.delete(`${API_URL}/delete-todo?id=${id}`);
      if (res.data.code == 200) {
        const newTodos = await getTodos();
        setTodos(newTodos);
        updateToggle(id);
      }
    } catch (err) {
      console.log(err);
    }
  };

  const save = async (todo: string) => {
    try {
      const res = await axios.post(`${API_URL}/create-todo`, {
        todo: todo,
      });
      if (res.data.code == 200) {
        const newTodos = await getTodos();
        setTodos(newTodos);
        setAddTodo(false);
      }
    } catch (err) {
      console.log(err);
    }
  };

  const search = () => {
    const t = todos.data.filter((item) =>
      item.todo.match(new RegExp(searchTodo.current.value))
    );
    setTodos({ data: t, code: todos.code, status: todos.status });
  };

  return (
    <>
      <div className="w-full flex flex-col h-full my-8 mx-4 md:mx-24">
        <div className="w-full">
          <div className="text-center font-bold mb-6 text-2xl text-blue">
            Todo List
          </div>
          <div className="max-w-4xl w-full mx-auto mb-6 flex flex-row">
            <div className="flex flex-row">
              {!addTodo && (
                <>
                  <input
                    ref={searchTodo}
                    placeholder="search todo"
                    className="bg-dark-primary border focus:outline-none border-red rounded-md mr-4 py-2 px-4"
                  />
                  <button
                    onClick={search}
                    className="bg-green px-4 text-dark-primary py-2 rounded-md"
                  >
                    search
                  </button>
                </>
              )}
            </div>

            <div className="flex w-full justify-end">
              {!addTodo && (
                <button
                  onClick={() => setAddTodo(true)}
                  className="bg-purple px-4 text-dark-primary py-2 rounded-md"
                >
                  add todo
                </button>
              )}
              {addTodo && (
                <div className="flex flex-row w-full">
                  <input
                    ref={dataAddTodo}
                    placeholder="todo"
                    className="bg-dark-primary w-full border focus:outline-none border-red rounded-md mr-4 py-2 px-4"
                  />
                  <button
                    onClick={() => save(dataAddTodo.current.value)}
                    className="bg-purple mr-4 px-4 text-dark-primary py-2 rounded-md"
                  >
                    save
                  </button>
                  <button
                    onClick={() => setAddTodo(false)}
                    className="bg-red px-4 text-dark-primary py-2 rounded-md"
                  >
                    cancel
                  </button>
                </div>
              )}
            </div>
          </div>
          <div>
            <table className="mx-auto max-w-4xl w-full whitespace-nowrap rounded-lg bg-dark divide-y divide-gray-300 overflow-hidden">
              <thead className="border-b-4 bg-aqua text-dark-secondary border-blue">
                <tr className="text-left">
                  <th className="px-4 w-1/12 text-center"> no </th>
                  <th className="px-4"> todo </th>
                  <th className="px-4 py-2 text-center w-1/4"> action </th>
                </tr>
              </thead>
              <tbody className="bg-dark-secondary divide-y divide-white divide-opacity-5">
                {todos.data.map((todo, index) => (
                  <tr key={index} className="">
                    <td className="text-center">
                      <span className="w-1/12">{index + 1}</span>
                    </td>
                    <td>
                      {todoUpdate.id !== todo.id && (
                        <div className="w-full px-4 ">{todo.todo}</div>
                      )}
                      {todoUpdate.id == todo.id && (
                        <div>
                          <input
                            type="text"
                            placeholder="todo"
                            className="bg-dark-primary py-1 border rounded-md px-4 focus:outline-none border-purple"
                            defaultValue={todo.todo}
                            onChange={(e) => setTodo(e.target.value)}
                          />
                        </div>
                      )}
                    </td>
                    <td className="px-6 py-4 text-center">
                      {todoUpdate.id != todo.id && (
                        <div>
                          <button
                            onClick={() => updateToggle(todo.id)}
                            className="bg-yellow rounded-md px-4 py-1 mr-4 text-dark-secondary"
                          >
                            update
                          </button>
                          <button
                            onClick={() => remove(todo.id)}
                            className="bg-red px-4 py-1 text-dark-primary rounded-md"
                          >
                            delete
                          </button>
                        </div>
                      )}
                      {todoUpdate.id == todo.id && (
                        <div>
                          <button
                            onClick={() => edit(todo.id)}
                            className="bg-green px-4 py-1 mr-4 text-dark-primary rounded-md"
                          >
                            save
                          </button>
                          <button
                            onClick={() => updateToggle(todo.id)}
                            className="bg-purple px-4 py-1 text-dark-primary rounded-md"
                          >
                            cancel
                          </button>
                        </div>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </>
  );
}
