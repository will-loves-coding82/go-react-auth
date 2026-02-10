import { Link } from "react-router";
import { useAuth } from "../context/AuthProvider";


export function Login() {

    const { isAuthenticated, isFetchingUser, logIn } = useAuth();
    if (isFetchingUser) return null;
    return (
        <section className="h-screen bg-black flex flex-col gap-4 justify-center">
            <span className="mx-auto">
                {isAuthenticated && <div id="status" className="bg-green-500 text-sm font-bold text-green-200 rounded-lg py-1 px-2 ">authorized</div>}
                {!isAuthenticated && <div id="status" className="w-fit bg-orange-500 text-sm font-bold text-orange-200 rounded-lg py-1 px-2 ">unauthorized</div>}
            </span>
    
            <header className="mx-auto text-center" >
                <h1 className="text-4xl text-white font-semibold">Go Auth + React</h1>
                <p className="text-slate-400 max-w-sm mt-2">This is an example application that interacts with Open Authorization workflows managed by a dedicated Go server.</p>

                { !isAuthenticated ?
                    <button className="mt-8 border-1 border-slate-400 px-2 py-1 rounded-sm hover:cursor-pointer hover:bg-slate-800" >
                        <p className="text-white text-base font-medium" onClick={logIn}>Log In</p>
                    </button>
                    :
                    <button  className="mt-8 border-1 border-slate-800 px-2 py-1 rounded-sm hover:cursor-pointer hover:bg-slate-800" >
                        <Link to="/dashboard" className="text-white text-base font-medium">Dashboard</Link>
                    </button>
                }
            </header>
        </section>
    )
}