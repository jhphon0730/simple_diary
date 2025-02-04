import React from 'react'

import Sidebar from '@/components/Sidebar'
import Navbar from '@/components/Navbar'
import WithAuthLayout from '@/components/config/WithAuthLayout';
import { SidebarProvider } from '@/components/ui/sidebar'

type MainLayoutProps = {
	children: React.ReactNode;
};

const MainLayout = ({ children }: MainLayoutProps) => {
	return (
		<WithAuthLayout>
			<SidebarProvider>
				<div className="flex w-screen h-screen">
					<Sidebar />
					<div className="flex flex-col flex-grow">
						<Navbar />
						<main className="flex-grow p-6 overflow-auto">
							{children}
						</main>
					</div>
				</div>
			</SidebarProvider>
		</WithAuthLayout>
	);
};

export default MainLayout;
