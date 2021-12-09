import axios from 'axios';
import { useEffect, useState } from 'react';
import { API_URL } from '../../constants';

export function useGetTodos() {
  const [todos, setTodos] = useState<ApiRes>(null);

  useEffect(() => {
    getTodos()
      .then((res) => {
        setTodos(res);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return { todos, setTodos };
}

export async function getTodos() {
  try {
    const res = await axios.get(`${API_URL}/todo-list`);
    return Promise.resolve(res.data);
  } catch (err) {
    return Promise.reject(err.message);
  }
}

type ApiRes = {
  status: string;
  code: number;
  data: Array<Todo>;
};

export type Todo = {
  id: number;
  todo: string;
};
