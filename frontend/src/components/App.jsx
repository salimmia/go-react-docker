import Home from "./Home";
import Login from "./Login";
import Register from "./Register";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import UpdateUser from "./UpdateUser";
import Navbar from "./Navbar";
import UserProfile from "./UserProfile";

function App() {
  return (
    <div style={{ marginTop: "-3.5rem" }}>
      <BrowserRouter>
        <Navbar>
          <Login />
        </Navbar>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/users/register" element={<Register />} />
          <Route path="/users/login" element={<Login />} />
          <Route
            path="/users/profile/update-user/:userId"
            element={<UpdateUser />}
          />
          <Route path="/users/profile/:userId" element={<UserProfile />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
