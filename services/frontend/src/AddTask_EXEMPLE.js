import React from "react";
import { useMutation } from "@apollo/react-hooks";
import { ADD_TASK } from "./Query";
const AddTask = ({ refetch }) => {
  let input;
  let completed;
  const [addTasks, { data }] = useMutation(ADD_TASK);
  console.log(data);
  return (
    <div>
      <form
        onSubmit={e => {
          e.preventDefault();
          addTasks({
            variables: {
              input: {
                name: input.value,
                completed: completed.checked
              }
            }
          });
          input.value = "";
          completed.checked = false;
          refetch();
        }}
      >
        <input
          ref={node => {
            input = node;
          }}
        />
        <input
          type="checkbox"
          ref={node => {
            completed = node;
          }}
        />
        <button type="submit">Add Task</button>
      </form>
    </div>
  );
};

export { AddTask };
