import { Outlet } from "react-router-dom";
import { Toaster } from "@/components/ui/toaster";
const MainLayout = () => {
  return (
    <>
      <main className="container mx-auto sm:px-6 lg:px-8">
        <Outlet />
      </main>
      <Toaster />
    </>
  );
};

export default MainLayout;
