def prompt(question, default=None):
    if default is None:
        default_text = ''
    else:
        default_text = ' [{0}]'.format(default)
    question = '{0}{1}? '.format(question, default_text)
    answer = input(question).rstrip('\n')
    if not answer:
        answer = default
    return answer
