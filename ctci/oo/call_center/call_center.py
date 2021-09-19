RESPONDENT = "respondent"
MANAGER = "manager"
DIRECTOR = "director"


class Employee:
    def __init__(self, _type, name="", done_callback=None):
        self.name = name
        self.busy = False
        self.type = _type
        self.done_callback = done_callback

    def __repr__(self) -> str:
        busy_str = "busy" if self.busy else "free"
        return f"{self.name} - {self.type} - ({busy_str})"

    def handle_call(self):
        self.busy = True

    def finish_call(self):
        self.busy = False
        if self.done_callback:
            self.done_callback(self)


class Respondent(Employee):
    def __init__(self, *args, **kwargs):
        super().__init__("respondent", *args, **kwargs)


class Manager(Employee):
    def __init__(self, *args, **kwargs):
        super().__init__("manager", *args, **kwargs)


class Director(Employee):
    def __init__(self, *args, **kwargs):
        super().__init__("director", *args, **kwargs)


class Roster:
    def __init__(self) -> None:
        self.employees = []
        self.call_router = CallRouter(self.employees)

    def __repr__(self) -> str:
        return "\n".join([str(e) for e in self.employees])

    def add_respondent(self, name):
        self.__add_employee(Respondent(
            name, done_callback=self.call_router.add))

    def add_manager(self, name):
        self.__add_employee(Manager(name, done_callback=self.call_router.add))

    def add_director(self, name):
        self.__add_employee(Director(name, done_callback=self.call_router.add))

    def __add_employee(self, employee: Employee):
        self.employees.append(employee)
        self.call_router.add(employee)

    def remove_employee(self, employee: Employee):
        self.employees.remove(employee)
        self.call_router.remove(employee)


class CallRouter:
    def __init__(self, employees) -> None:
        # capture the ordering of the escalation tier in the list.
        self.tiers = [RESPONDENT, MANAGER, DIRECTOR]
        self.free_employees = {tier: [] for tier in self.tiers}
        for e in employees:
            self.add(e)

    def add(self, employee: Employee):
        self.free_employees[employee.type].append(employee)

    def remove(self, employee: Employee):
        self.free_employees[employee.type].remove(employee)

    def dispatch_call(self):
        for tier in self.tiers:
            try:
                employee = self.free_employees[tier].pop()
                employee.handle_call()
                return
            except IndexError:
                continue
        raise Exception("No free employees to handle call.")


if __name__ == "__main__":
    r = Roster()
    r.add_respondent("jim")
    r.add_respondent("sally")
    r.add_manager("bob")
    r.add_director("jen")
    print(r)
    r.call_router.dispatch_call()
    print("After call 1")
    print(r)
    r.call_router.dispatch_call()
    print("After call 2")
    print(r)
    r.call_router.dispatch_call()
    print("After call 3")
    print(r)
    r.call_router.dispatch_call()
    print("After call 4")
    print(r)
    r.call_router.dispatch_call()
    print("After call 5")
    print(r)
