RESPONDENT = "respondent"
MANAGER = "manager"
DIRECTOR = "director"


class Employee:
    def __init__(self, _type, name=""):
        self.name = name
        self.busy = False
        self.type = _type

    def __repr__(self) -> str:
        busy_str = "busy" if self.busy else "free"
        return f"{self.name} - {self.type} - ({busy_str})"

    def handle_call(self):
        self.busy = True

    def finish_call(self):
        self.busy = False


class Respondent(Employee):
    def __init__(self):
        super().__init__("respondent")


class Manger(Employee):
    def __init__(self):
        super().__init__("manager")


class Director(Employee):
    def __init__(self):
        super().__init__("director")


class Roster:
    def __init__(self) -> None:
        self.employees = []
        self.call_router = CallRouter(self.employees)

    def __repr__(self) -> str:
        return "\n".join(self.employees)

    def add_employee(self, employee: Employee):
        self.employees.append(employee)
        self.call_router.add(employee)

    def remove_employee(self, employee: Employee):
        self.employees.remove(employee)
        self.call_router.remove(employee)


class CallRouter:
    def __init__(self) -> None:
        # capture the ordering of the escalation tier in the list.
        self.tiers = [RESPONDENT, MANAGER, DIRECTOR]
        self.free_employees = {tier: [] for tier in self.tiers}

    def add(self, employee: Employee):
        self.free_employees[employee.type].append(employee)

    def remove(self, employee: Employee):
        self.free_employees[employee.type].remove(employee)

    def dispatch_call(self):
        for tier in self.tiers:
            try:
                employee = self.free_employees[tier].pop()
                employee.handle_call()
            except IndexError:
                continue
        raise Exception("No free employees to handle call.")
