import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: '다이어리',
  description: '커플을 위한 공유 다이어리 앱',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ko">
      <body className={inter.className}>
				<main className='min-h-screen min-w-full'>
          {children}
				</main>
      </body>
    </html>
  )
}

