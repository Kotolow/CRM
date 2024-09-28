import React from 'react';
import {GrNotification} from "react-icons/gr";
import { FiSearch, FiPlus } from 'react-icons/fi';


export default function Header({ title, buttonText }) {
    return (
        <div>
        <div className="top-bar ">
            <a className="notifications flex items-center text-gray-700 hover:text-blue-500 p-2 bg-white hover:bg-gray-200 rounded-lg mr-4 cursor-pointer">
                <GrNotification className="w-6 h-6"/>
            </a>
            <a href='/profile'
                className="user-info flex items-center text-gray-700 hover:text-blue-500 p-2 bg-white hover:bg-gray-200 rounded-lg m-0 cursor-pointer">
                <img src="profile_icon.png" alt="Profile" className="w-8 h-8 rounded-full mr-2"/>
                <span>Evan Yates</span>
            </a>
            </div>
            <div className="lower-bar flex justify-between items-center mb-6 mt-5">
                <div>
                    <h1 className="text-3xl font-bold text-gray-800">{title}</h1>
                </div>
            <div className="flex items-center space-x-4">
                <div className="relative">
                    <input
                        type="text"
                        placeholder="Search"
                        className="border rounded-lg p-2 pl-10 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                    <FiSearch className="absolute left-2 top-2 text-gray-400" />
                </div>
            <button className="bg-blue-500 text-white px-4 py-2 rounded-lg flex items-center hover:bg-blue-600">
            <FiPlus className="mr-2" />
           {buttonText}
          </button>
        </div>
    </div>
    </div>
    );
}
