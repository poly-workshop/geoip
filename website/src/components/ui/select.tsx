import * as React from "react"
import { cn } from "@/lib/utils"

interface SelectProps {
  value?: string
  onValueChange?: (value: string) => void
  children: React.ReactNode
}

interface SelectTriggerProps extends React.HTMLAttributes<HTMLButtonElement> {
  className?: string
}

interface SelectValueProps {
  placeholder?: string
}

interface SelectContentProps {
  children: React.ReactNode
}

interface SelectItemProps {
  value: string
  children: React.ReactNode
}

const Select: React.FC<SelectProps> = ({ value, onValueChange, children }) => {
  const [isOpen, setIsOpen] = React.useState(false)
  
  const handleValueChange = (newValue: string) => {
    onValueChange?.(newValue)
    setIsOpen(false)
  }

  return (
    <div className="relative">
      {React.Children.map(children, child => 
        React.isValidElement(child) 
          ? React.cloneElement(child as React.ReactElement<Record<string, unknown>>, {
              value: value,
              onValueChange: handleValueChange,
              isOpen,
              setIsOpen
            })
          : child
      )}
    </div>
  )
}

const SelectTrigger = React.forwardRef<HTMLButtonElement, SelectTriggerProps & {
  isOpen?: boolean
  setIsOpen?: (open: boolean) => void
}>(({ className, children, isOpen, setIsOpen, ...props }, ref) => (
  <button
    ref={ref}
    type="button"
    role="combobox"
    aria-expanded={isOpen}
    className={cn(
      "flex h-10 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
      className
    )}
    onClick={() => setIsOpen?.(!isOpen)}
    {...props}
  >
    {children}
    <svg
      width="15"
      height="15"
      viewBox="0 0 15 15"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={cn("h-4 w-4 opacity-50 transition-transform", isOpen && "rotate-180")}
    >
      <path
        d="m4.5 6 3 3 3-3"
        stroke="currentColor"
        strokeWidth="1"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  </button>
))
SelectTrigger.displayName = "SelectTrigger"

const SelectValue: React.FC<SelectValueProps & { value?: string }> = ({ placeholder, value }) => (
  <span className={cn(!value && "text-muted-foreground")}>
    {value || placeholder}
  </span>
)

const SelectContent: React.FC<SelectContentProps & {
  isOpen?: boolean
  onValueChange?: (value: string) => void
  value?: string
}> = ({ children, isOpen, onValueChange, value }) => {
  if (!isOpen) return null
  
  return (
    <div className="absolute top-full z-50 min-w-32 w-full overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-md animate-in fade-in-0 zoom-in-95">
      <div className="p-1">
        {React.Children.map(children, child =>
          React.isValidElement(child)
            ? React.cloneElement(child as React.ReactElement<Record<string, unknown>>, { 
                onValueChange,
                selectedValue: value
              })
            : child
        )}
      </div>
    </div>
  )
}

const SelectItem: React.FC<SelectItemProps & { 
  onValueChange?: (value: string) => void
  selectedValue?: string
}> = ({
  value,
  children,
  onValueChange,
  selectedValue
}) => {
  const isSelected = selectedValue === value
  
  return (
    <div
      className={cn(
        "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
        isSelected && "bg-accent text-accent-foreground"
      )}
      onClick={() => onValueChange?.(value)}
    >
      {children}
      {isSelected && (
        <svg
          width="15"
          height="15"
          viewBox="0 0 15 15"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
          className="ml-auto h-4 w-4"
        >
          <path
            d="m4.5 7.5 2 2 4-4"
            stroke="currentColor"
            strokeWidth="1"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      )}
    </div>
  )
}

export {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
}