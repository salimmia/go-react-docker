import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const Home = () => {
  const [userData, setUserData] = useState([]);

  useEffect(() => {
    fetchUserData();
  }, []);

  const fetchUserData = async () => {
    try {
      const response = await axios.get("http://localhost:8080/users");
      console.log(response.data);
      setUserData(response.data);
    } catch (error) {
      console.error("Error fetching user information:", error);
    }
  };

  const getAllUsers = () => {
    fetchUserData();
  };

  return (
    <div className="container">
      <header>
        <h1>Login Success Page</h1>
        <Link to="/users/login">Logout</Link>
      </header>
      <main>
        <button onClick={getAllUsers}>Get All Users</button>
        <h2>User Information</h2>
        <div className="row">
          {userData.map((user) => (
            <div key={user.id} className="col-md-6">
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">
                    {user.first_name} {user.last_name}
                  </h5>
                  <p className="card-text">
                    <strong>Email:</strong> {user.email}
                  </p>
                  <p className="card-text">
                    <strong>Phone:</strong>{" "}
                    {user.profile?.phone_number || "N/A"}
                  </p>
                  <p className="card-text">
                    <strong>Birth Date:</strong>{" "}
                    {user.profile?.birth_date
                      ? user.profile.birth_date.substring(0, 10)
                      : "N/A"}
                  </p>
                  <p className="card-text">
                    <strong>User Name:</strong> {user.thumbnail}
                  </p>
                  {/* Add a link to the user's profile */}
                  <Link to={`/users/profile/${user.id}`}>View Profile</Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      </main>
    </div>
  );
};

export default Home;
