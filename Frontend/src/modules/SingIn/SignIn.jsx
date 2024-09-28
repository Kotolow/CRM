import React from 'react';

function SignIn() {  return (
    <div className="flex h-screen items-center justify-center bg-gray-100">
      <div className="flex shadow-lg bg-white rounded-lg overflow-hidden w-2/3 h-3/4">

        <div className="w-1/2 bg-blue-500 flex items-center justify-center p-10">
          <div className="text-white">
            <img
              src="https://img.icons8.com/ios-filled/50/ffffff/check-file.png"
              alt="Logo"
              className="mb-6"
            />
            <h2 className="text-5xl font-bold mb-4">H&H CRM</h2>
            <p className="text-xl">Your place to work</p>
            <p className="text-xl font-semibold">Plan. Create. Control.</p>
    
          </div>
        </div>

        {/* Правая часть с формой */}
        <div className="w-1/2 p-8 flex flex-col justify-center">
          <h2 className="text-3xl font-bold mb-6">Sign In to H&H</h2>
          <form>
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
            <div className="mb-6">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Password
              </label>
              <input
                type="password"
                placeholder="********"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
            </div>
            <div className="flex items-center justify-between mb-6">
              <label className="flex items-center">
                <input type="checkbox" className="mr-2" />
                Remember me
              </label>
              {/* <a href="#" className="text-sm text-blue-500 hover:underline">
                Forgot Password?
              </a> */}
            </div>
            <button
              type="submit"
              className="w-full bg-blue-500 text-white p-3 rounded-lg font-bold hover:bg-blue-600 transition duration-300"
            >
              Sign In
            </button>
          </form>
          <p className="mt-4 text-center text-gray-600">
            Don't have an account?{' '}
            <a href="/register" className="text-blue-500 hover:underline">
              Sign up
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}

export default SignIn;
