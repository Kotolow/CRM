// src/components/TaskCard.js
import React from 'react';

const TaskCard = ({ title, time, users }) => {
  return (
    <div className="bg-white p-4 rounded-lg shadow-lg mb-4">
      <h3 className="text-lg font-semibold mb-2">{title}</h3>
      <p className="text-gray-500 text-sm mb-2">{time}</p>
      <div className="flex -space-x-2">
        {users.map((user, index) => (
          <img
            key={index}
            src={user}
            alt="User"
            className="w-8 h-8 rounded-full border-2 border-white"
          />
        ))}
      </div>
    </div>
  );
};

export default TaskCard;
