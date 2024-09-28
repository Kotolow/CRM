import React, { useState, useEffect } from 'react';
import { DndProvider, useDrag, useDrop } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import Sidebar from '../Sidebar/Sidebar';
import Header from "../Header/Header";
import '../Projects/Projects.scss'
import axios from 'axios';



const initialData = {
  tasks: {
    'task-1': { id: 'task-1', title: 'UX sketches', time: '4d', users: ['https://via.placeholder.com/32'] },
    'task-2': { id: 'task-2', title: 'Mind Map', time: '2d', users: ['https://via.placeholder.com/32'] },
    'task-3': { id: 'task-3', title: 'Research reports', time: '4d', users: ['https://via.placeholder.com/32'] },
  },
  columns: {
    'column-1': {
      id: 'open',
      title: 'Open',
      taskIds: ['task-1', 'task-2'],
    },
    'column-2': {
      id: 'column-2',
      title: 'In Progress',
      taskIds: [],
    },
    'column-3': {
      id: 'column-3',
      title: 'In Review',
      taskIds: ['task-3'],
    },
    'column-4': {
      id: 'column-4',
      title: 'Done',
      taskIds: [],
    },
  },
  columnOrder: ['column-1', 'column-2', 'column-3', 'column-4'],
};

const ItemTypes = {
  TASK: 'task',
};

const TaskCard = ({ task, index }) => {
  const [, ref] = useDrag({
    type: ItemTypes.TASK,
    item: { id: task.id, index },
  });

  return (
    <div ref={ref} className="bg-white p-4 rounded-lg shadow mb-2">
      <h4>{task.title}</h4>
      <p>{task.time}</p>
    </div>
  );
};


const Column = ({ column, tasks, moveTask }) => {
  const [, drop] = useDrop({
    accept: ItemTypes.TASK,
    drop: (item) => moveTask(item.id, column.id),
  });

  return (
    <div ref={drop} className="bg-gray-200 p-4 rounded-lg min-h-[400px]">
      <h2 className="text-xl font-semibold mb-4">{column.title}</h2>
      {tasks.map((task, index) => (
        <TaskCard key={task.id} task={task} index={index} />
      ))}
    </div>
  );
};

const Dashboard = () => {
  const [data, setData] = useState(initialData);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await axios.get('https://69fbdebe0a2299.lhr.life/v1/projects/H%26h/tasks');
        const tasks = response.data.tasks;
        console.log(tasks);
      } catch (error) {
        console.error('Ошибка при загрузке задач:', error);
      }
    };

    fetchTasks();
  }, []);

  const moveTask = (taskId, destinationColumnId) => {
    const task = data.tasks[taskId];

    if (!task) {
      console.error('Task not found:', taskId);
      return;
    }

    const sourceColumnId = Object.keys(data.columns).find((columnId) =>
      data.columns[columnId].taskIds.includes(taskId)
    );

    if (!sourceColumnId || !destinationColumnId) {
      console.error('Source or destination column not found:', sourceColumnId, destinationColumnId);
      return;
    }

    const sourceColumn = data.columns[sourceColumnId];
    const destinationColumn = data.columns[destinationColumnId];

    if (sourceColumnId === destinationColumnId) {
      return;
    }

    const newSourceTaskIds = sourceColumn.taskIds.filter((id) => id !== taskId);

    const newDestinationTaskIds = [...destinationColumn.taskIds, taskId];

    const newColumns = {
      ...data.columns,
      [sourceColumnId]: { ...sourceColumn, taskIds: newSourceTaskIds },
      [destinationColumnId]: { ...destinationColumn, taskIds: newDestinationTaskIds },
    };

    setData({ ...data, columns: newColumns });
  };

  return (
    <DndProvider backend={HTML5Backend}>
      <div className="flex h-screen bg-gray-100">
        <Sidebar />
        <div className='content'>  
          <Header title="Dashboard" buttonText="Add Task"/>
        <div className="flex-1">

          <div className="grid grid-cols-4 gap-6">
            {data.columnOrder.map((columnId) => {
              const column = data.columns[columnId];
              const tasks = column.taskIds.map((taskId) => data.tasks[taskId]);

              return (
                <Column
                  key={column.id}
                  column={column}
                  tasks={tasks}
                  moveTask={moveTask}
                />
              );
            })}
          </div>
        </div>
      </div>
      </div>
    </DndProvider>
  );
};

export default Dashboard;
