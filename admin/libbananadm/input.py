import sys


def prompt(question, default=None):
    if default is None:
        default_text = ''
    else:
        default_text = ' [{0}]'.format(default)
    question = '{0}{1}? '.format(question, default_text)
    sys.stderr.write(question)
    answer = input().rstrip('\n')
    if not answer:
        answer = default
    return answer
