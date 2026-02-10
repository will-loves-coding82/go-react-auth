import { Navigate, useLocation, Outlet} from "react-router-dom";
import { useAuth } from "../context/AuthProvider";

export const ProtectedRoute = () => {
    const {isAuthenticated, isFetchingUser} = useAuth();
    let location = useLocation();
    // while auth is being fetched, don't redirect — render nothing or a loader
    if (isFetchingUser) return null;
    return isAuthenticated ? <Outlet/> : <Navigate to="/" replace state={{ from: location }} />;
};
