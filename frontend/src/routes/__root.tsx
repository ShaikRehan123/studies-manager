import * as React from 'react'
import { Link, Outlet, createRootRoute } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  return (
    <div className='dark:bg-gray-600 bg-gray-200 text-white min-h-screen p-2'>
      <Outlet />
      <TanStackRouterDevtools position="bottom-right" />
    </div>
  )
}
