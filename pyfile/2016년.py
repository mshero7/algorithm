def solution(a, b):
    month = [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    date = ['FRI','SAT','SUN','MON','TUE','WED','THU']

    if a > 1:
        answer = date[(sum(month[:a-1]) + b) % 7 - 1]
    else:
        answer = date[b%7-1]
    return answer

print(solution(5,24))
