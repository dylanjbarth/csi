#include <string.h>
#include <stdio.h>
#include "grade_school.h"

static roster_t roster = {0};

roster_t get_roster()
{
  return roster;
}

void clear_roster()
{
  roster.count = 0;
  memset(roster.students, 0, MAX_STUDENTS);
}

roster_t get_grade(uint8_t desired_grade)
{
  roster_t grade_x = {0};
  for (size_t i = 0; i < roster.count; i++)
  {
    if (roster.students[i].grade == desired_grade)
    {
      grade_x.students[grade_x.count] = roster.students[i];
      grade_x.count += 1;
    }
  }
  return grade_x;
}

int add_student(char *name, uint8_t grade)
{
  print_roster();
  printf("Add student called with grade: %d and name: %s\n", grade, name);
  // insert in sorted order
  int inserted = 0;
  for (size_t i = 0; i < roster.count; i++)
  {
    student_t cur_student = roster.students[i];
    printf("Iterating through students at pos %lu, current roster size is %lu\n", i, roster.count);
    int lesser_grade = grade < cur_student.grade;
    int same_grade = grade == cur_student.grade;
    int lesser_alpha = strcmp(name, cur_student.name) < 0;
    printf("Lesser grade %d; Same grade %d; lesser alpha %d\n", lesser_grade, same_grade, lesser_alpha);
    if (lesser_grade || (same_grade && lesser_alpha))
    {
      shift_students(i);
      roster.students[i] = (student_t){grade, name};
      roster.count++;
      inserted = 1;
      break;
    }
  }
  if (!inserted)
  {
    roster.students[roster.count] = (student_t){grade, name};
    roster.count++;
  }
  print_roster();
  return 1;
}

// shifts students starting at position ps
void shift_students(size_t pos)
{
  printf("Shifting students from pos %lu\n", pos);
  for (size_t i = roster.count; i > pos; i--)
  {
    roster.students[i] = roster.students[i - 1];
  }
}

void print_roster()
{
  printf("\n\n***Current roster:***\n\n");
  for (size_t i = 0; i < roster.count; i++)
  {
    printf("Grade: %d; Name: %s\n", roster.students[i].grade, roster.students[i].name);
  }
  printf("\n\n***End roster:***\n");
}
