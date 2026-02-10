import { useQuery } from '@tanstack/react-query'

function AthleteList() {
  const { data: athletes, isLoading, error } = useQuery({
    queryKey: ['athletes'],
    queryFn: async () => {
      const response = await fetch('/api/athletes')
      if (!response.ok) {
        throw new Error('Failed to fetch athletes')
      }
      return response.json()
    },
  })

  if (isLoading) {
    return (
      <div className="text-center py-8">
        <p className="text-gray-500 animate-pulse text-lg">Loading athletes...</p>
      </div>
    )
  }

  if (error) {
    return (
      <div className="bg-red-50 border border-red-200 rounded-lg p-4 text-center">
        <p className="text-red-700">Error: {error.message}</p>
      </div>
    )
  }

  return (
    <div className="w-full max-w-2xl">
      <h2 className="text-2xl font-bold text-gray-900 mb-4">Athletes</h2>
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Grade</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Personal Record</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {athletes.map((athlete) => (
              <tr key={athlete.id}>
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{athlete.name}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{athlete.grade}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{athlete.personalRecord}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default AthleteList
