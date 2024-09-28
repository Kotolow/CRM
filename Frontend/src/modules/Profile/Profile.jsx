import React from 'react';
import Sidebar from '../Sidebar/Sidebar';

function ProfilePage() {
  return (
    <div className="flex h-screen bg-gray-100 content">
      <Sidebar />
      {/* User Profile Section */}
      <div className="w-1/4 bg-white p-6 flex flex-col items-center shadow-lg m-6 rounded-lg">
        <img
          src="https://via.placeholder.com/150"
          alt="User profile"
          className="w-32 h-32 rounded-full mb-4"
        />
        <h2 className="text-lg font-semibold">Evan Yates</h2>
        <div className="mt-4">
          <label className="align-start text-gray-600">Email</label>
          <p>evanyates@gmail.com</p>

          <label className="text-gray-600">Location</label>
          <p>NYC, New York, USA</p>

          <label className="text-gray-600">Mobile</label>
          <p>+1 675 342-21-10</p>

          <label className="text-gray-600">Skype</label>
          <p>Evan 2216</p>
        </div>
      </div>

      {/* Main Content */}
      <div className="flex-1 p-10">
        {/* Tabs */}
    

        {/* Content Area */}
        <div className="bg-white p-6 rounded-lg shadow">

          {/* Achievements Section (Example content) */}
          <div className="mt-10">
            <h3 className="text-xl font-semibold mb-4">Achievements</h3>
            <ul className="list-disc pl-6">
              <li>Completed 100+ projects</li>
              <li>Top designer award (2023)</li>
              <li>Certified in UX/UI design</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ProfilePage;
