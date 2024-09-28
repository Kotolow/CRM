// src/Register.js
import React from 'react';

function Register() {
  return (
    <div className="flex h-screen items-center justify-center bg-gray-100">
      <div className="flex shadow-lg bg-white rounded-lg overflow-hidden w-2/3 h-3/4">
        {/* Левая часть с изображением */}
        <div className="w-1/3 bg-blue-500 flex items-center justify-center p-10">
          <div className="text-white">
            <img
              src="https://img.icons8.com/ios-filled/50/ffffff/check-file.png"
              alt="Logo"
              className="mb-6"
            />
            <h2 className="text-4xl font-bold mb-4">H&H CRM</h2>
            <p className="text-lg">Get started</p>
          </div>
        </div>

        {/* Правая часть с формой регистрации */}
        <div className="w-1/2 p-8 flex flex-col justify-center">
          <h2 className="text-3xl font-bold mb-6">Sign Up to H&H</h2>
          <form>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Full Name
              </label>
              <input
                type="text"
                placeholder="John Doe"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Email Address
              </label>
              <input
                type="email"
                placeholder="youremail@gmail.com"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Password
              </label>
              <input
                type="password"
                placeholder="********"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
            </div>
            <div className="mb-6">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Confirm Password
              </label>
              <input
                type="password"
                placeholder="********"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
            </div>
            <button
              type="submit"
              className="w-full bg-blue-500 text-white p-3 rounded-lg font-bold hover:bg-blue-600 transition duration-300"
            >
              Sign Up
            </button>
          </form>
          <p className="mt-4 text-center text-gray-600">
            Already have an account?{' '}
            <a href="/" className="text-blue-500 hover:underline">
              Sign In
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}

export default Register;
