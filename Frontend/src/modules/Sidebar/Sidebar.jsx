import React, {useState} from 'react';
import { FiLogOut } from 'react-icons/fi';
import { AiFillHome, AiFillProject } from 'react-icons/ai';
import { BsFillCalendarFill, BsFillPersonFill } from 'react-icons/bs';


const Sidebar = () => {

  const [active, setActive] = useState('projects')


  const toggleMenu = (item) => {
    setActive(item)
  }

  return (
    <div className="w-64 bg-white h-screen shadow-lg">
      <div className="p-6">
        <h1 className="text-2xl font-bold text-gray-800 mb-6">H&H CRM</h1>
        <ul>
          <li className='mb-6'>
            <a href="/projects" className={`flex items-center ${active === 'projects' ? 'text-blue-500' : 'text-gray-700'} hover:text-blue-500`}
               onClick={() => toggleMenu('projects')}>
              <AiFillHome className="mr-3"/>
              Projects
            </a>
          </li>
          <li className="mb-6">
            <a href="/dashboard" className={`flex items-center ${active === 'dashboard' ? 'text-blue-500' : 'text-gray-700'} hover:text-blue-500`}
               onClick={() => toggleMenu('dashboard')}>
              <AiFillProject className="mr-3" />
              Dashboard
            </a>
          </li>
          <li className="mb-6">
            <a href="/diagram" className={`flex items-center ${active === 'diagram' ? 'text-blue-500' : 'text-gray-700'} hover:text-blue-500`}
               onClick={() => toggleMenu('diagram')}>
              <BsFillCalendarFill className="mr-3" />
              Diagram
            </a>
          </li>
        </ul>
        <div className="absolute bottom-0 left-0 w-full p-6">
          <a href="/" className="flex items-center text-gray-700 hover:text-red-500">
            <FiLogOut className="mr-3" />
            Logout
          </a>
        </div>
      </div>
    </div>
    );
};

export default Sidebar;