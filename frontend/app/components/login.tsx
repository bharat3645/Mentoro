import { useState } from "react";
import { Form } from "@remix-run/react";

export default function Login() {
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  // Placeholder for form submission handler
  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setLoading(true);
    setError("");
    // TODO: Implement API call to backend for login
    // Example:
    // const res = await fetch("/api/v1/login", { ... })
    // if (!res.ok) setError("Invalid credentials");
    setTimeout(() => {
      setLoading(false);
      setError("This is a placeholder. Connect to backend API.");
    }, 1000);
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-purple-50">
      <div className="bg-white shadow-lg rounded-lg p-8 w-full max-w-md">
        <h1 className="text-2xl font-bold mb-6 text-center">Login</h1>
        <Form method="post" onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-1">
              Username or Email
            </label>
            <input
              type="text"
              id="username"
              name="username"
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
              Password
            </label>
            <input
              type="password"
              id="password"
              name="password"
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            />
          </div>
          {error && <div className="text-red-600 text-sm">{error}</div>}
          <button
            type="submit"
            className="w-full bg-primary-500 hover:bg-primary-600 text-white font-bold py-2 px-4 rounded-lg transition-all duration-200"
            disabled={loading}
          >
            {loading ? "Logging in..." : "Login"}
          </button>
        </Form>
        <div className="mt-4 text-center text-sm text-gray-500">
          Don&apos;t have an account? <a href="/register" className="text-primary-600 hover:underline">Register</a>
        </div>
      </div>
    </div>
  );
} 