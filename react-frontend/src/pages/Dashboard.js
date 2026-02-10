import { Link } from "react-router";
import { useAuth } from "../context/AuthProvider";
import { RxAvatar } from "react-icons/rx";
import { Navigate } from "react-router-dom";


export function Dashboard() {
    const { isAuthenticated, isFetchingUser, user, logout } = useAuth();

    return (
        <>
        {isFetchingUser && <p>loading...</p>}
        <section className="flex flex-col justify-center h-screen gap-8 bg-black h-screen text-white p-5 text-center items-center">

            {isAuthenticated && <div id="status" className="bg-green-500 text-sm font-bold text-green-200 rounded-lg py-1 px-2 ">authorized</div>}

            <h1 className="text-2xl text-white font-semibold">Dashboard</h1>
            <div className="m flex flex-col gap-4 mx-auto">
                {!user?.picture && <RxAvatar size={40} class="mx-auto"/> }
                {user?.picture &&                 
                    <img
                        src={user?.picture}
                        alt="google profile"
                        width={140}
                        referrerPolicy="no-referrer"
                        className="rounded-full mx-auto"
                    /> 
                }
                <p className="text-grey">{user?.email}</p>
            </div>

            <span className="flex gap-4 justify-center items-center">
                <Link to="/"  className="w-fit mx-auto bg-slate-800 px-2 py-1 rounded-sm text-white hover:cursor-pointer">Home</Link>
                <Link className="w-fit mx-auto bg-slate-200 px-2 py-1 rounded-sm text-black hover:cursor-pointer" onClick={logout}>Logout</Link>
            </span>

        </section>
        </>
    )
}