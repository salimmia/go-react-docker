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
    <div className="container-fluid p-0">
      <div
        style={{
          backgroundImage:
            "linear-gradient(#00d5ff,#0095ff,rgba(93,0,255,.555))",
        }}
        className="d-flex flex-column justify-content-center align-items-center text-center min-vh-100"
      >
        <h1 className="text-light">Login Success Page</h1>
        <Link to="/users/login" className="btn btn-light my-5">
          Logout
        </Link>
        <button className="btn btn-light my-1" onClick={getAllUsers}>
          Get All Users
        </button>
        <h1 className="text-light mb-4">User Information</h1>
        <div className="row">
          {userData.map((user) => (
            <div key={user.id} className="col-md-6">
              <div className="card mb-4">
                <div className="card-body">
                  <h5 className="card-title">
                    {user.first_name} {user.last_name}
                  </h5>
                  <p className="card-text">
                    <strong>Email:</strong> {user.email}
                  </p>
                  <p className="card-text">
                    <strong>Phone:</strong> {user.phone_number}
                  </p>
                  <p className="card-text">
                    <strong>Birth Date:</strong>{" "}
                    {user.profile.birth_date
                      ? user.profile.birth_date.substring(0, 10)
                      : ""}
                  </p>
                  <p className="card-text">
                    <strong>User Name:</strong> {user.thumbnail}
                  </p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Home;
