import { useContext, createContext } from "react";
import { useQuery, useQueryClient } from "@tanstack/react-query";

const AuthContext = createContext(undefined)

/**
 * fetchUser will make a GET request to the golang API to 
 * attempt to get the user. If an error is encountered, the function
 * returns null. Requests also attach credentials to ensure 
 * user-identifying browser cookies are processed by the server.
 */
const fetchUser = async () => {
  const response = await fetch("http://localhost:3000/user", { credentials: 'include' })
  if (!response.ok) { 
    return null 
  };
  const json = await response.json();
  return json.data ?? null;
}

export const AuthProvider = ({ children }) => {
  const queryClient = useQueryClient()
  const { data: user, isFetching: isFetchingUser } = useQuery({
    queryKey: ["auth"],
    queryFn: fetchUser,
    retry: false,
    refetchOnWindowFocus: false,
  })

  var isAuthenticated = !!user;

  /**
   * Performs a browser redirect to the golang login endpoint to begin
   * the OAuth flow with Google.
   */
  const logIn = () => {
    window.location.href = `http://localhost:3000/auth/login?provider=google`
  };

  /**
   * Performs a POST request to the golang logout endpoint to 
   * invalidate the user's session and logout.
   */
  const logout = async () => {
    try {
      const response = await fetch("http://localhost:3000/auth/logout", { method: 'POST', credentials: 'include' });
      if (response.status === 200) {
        queryClient.invalidateQueries({ queryKey: ["auth"] })
        window.location.href = "/"
      }
    }
    catch (error) {
      console.log(error)
    }
  };

  return (
    <AuthContext.Provider value={{ user, isAuthenticated, isFetchingUser, logIn, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (ctx === undefined) {
    throw new Error("useAuth must be used within AuthProvider");
  }
  return ctx;
};